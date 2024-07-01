package mongodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	mongodbModel "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	mongodb_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	log      *config.TerminalLog
}

type User struct {
	ID        string    `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	Role      string    `json:"Role"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func New(config *configuration.Configuration, logTerminal *config.TerminalLog) (*MongoDB, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", config.User, config.Password, config.Host, config.Port, config.DBName)
	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(options.Credential{
		Username: config.User,
		Password: config.Password,
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %w", err)
	}

	db := client.Database(config.DBName)

	return &MongoDB{
		Client:   client,
		Database: db,
		log:      logTerminal,
	}, nil
}

func (m *MongoDB) CreateUserStorage(userModel *domain.User, collectionName string) (string, error) {
	collection := m.Database.Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), userModel)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return "", err
	}

	return "User Created", nil
}

func (m *MongoDB) GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error) {
	collection := m.Database.Collection(collectionName)
	var result mongodbModel.UserModelMongoDB
	filter := bson.M{"email": userEmail}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, err
	}

	userModelDomain := &domain.User{
		ID:        result.ID.Hex(),
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		Role:      domain.UserRole(result.Role),
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return userModelDomain, nil
}

func (m *MongoDB) ListUsersStorage(collectionName string) ([]domain.User, error) {
	var userList []mongodbModel.UserModelMongoDB
	collection := m.Database.Collection(collectionName)

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}
	if err = cur.All(context.Background(), &userList); err != nil {
		m.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}

	var listUserAll []domain.User

	for _, user := range userList {
		listUserAll = append(listUserAll, domain.User{
			ID:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Password:  "",
			Role:      domain.UserRole(user.Role),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})

	}

	return listUserAll, nil
}

func (m *MongoDB) GetUserStorage(idUser string, collectionName string) (*domain.User, error) {

	var userBody domain.User
	objectId, err := primitive.ObjectIDFromHex(idUser)
	filter := bson.M{"_id": objectId}
	collection := m.Database.Collection(collectionName)
	err = collection.FindOne(context.Background(), filter).Decode(&userBody)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}

	user := &domain.User{
		ID:        userBody.ID,
		Name:      userBody.Name,
		Email:     userBody.Email,
		Role:      userBody.Role,
		CreatedAt: userBody.CreatedAt,
		UpdatedAt: userBody.UpdatedAt,
	}

	return user, nil
}

func (m *MongoDB) CreateNumberGiftCardStorage(amount int, collectionName string, infoToken *domain.TokenPayload, infoCoupon *domain.Coupon) ([]mongodb_model.CodeCoupon, error) {
	var codeCoupon []mongodbModel.CodeCoupon
	collection := m.Database.Collection(collectionName)

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}
	if err = cur.All(context.Background(), &codeCoupon); err != nil {
		m.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}

	var listCode []string
	var allCoupons []mongodb_model.CodeCoupon

	listCode = util.RandSeq(8, amount)

	for i := 0; i < len(listCode); i++ {
		exist, _ := util.SearchCode(listCode[i], codeCoupon)
		if exist {
			retryCode := util.RandSeq(8, 1)
			listCode = append(listCode, retryCode[0])
			i--
			continue
		}

		codeOne := &mongodbModel.CodeCoupon{
			ID:                primitive.NewObjectID(),
			UserOwner:         infoToken.UserID,
			Code:              listCode[i],
			IsUsed:            false,
			Price:             infoCoupon.PriceCoupon,
			IDReferentProcess: infoCoupon.IDReference,
		}

		_, err = collection.InsertOne(context.Background(), codeOne)
		if err != nil {
			m.log.ErrorLog.Println(err)
			return nil, err
		}

		allCoupons = append(allCoupons, mongodb_model.CodeCoupon{
			ID:        codeOne.ID,
			UserOwner: infoToken.UserID,
			Code:      listCode[i],
			IsUsed:    false,
			Price:     infoCoupon.PriceCoupon,
		})
	}

	return allCoupons, nil
}

func (m *MongoDB) LinkingGiftCardUserStorage(collectionName string, coupons []mongodb_model.CodeCoupon, infoPayment *domain.ResponseUalabisPOST, infoDomainCoupon *domain.Coupon) (*mongodb_model.CouponModel, error) {

	collection := m.Database.Collection(collectionName)

	var couponMetaData []mongodb_model.CouponMetaData

	for j := 0; j < len(coupons); j++ {
		couponMetaData = append(couponMetaData, mongodb_model.CouponMetaData{
			Code:   coupons[j].ID,
			IdLink: infoDomainCoupon.IDReference,
		})
	}

	modelCoupon := &mongodb_model.CouponModel{
		IDReferentProcess: infoDomainCoupon.IDReference,
		Email:             infoDomainCoupon.Email,
		Owner:             infoDomainCoupon.Owner,
		Title:             infoDomainCoupon.Title,
		Description:       infoDomainCoupon.Description,
		AmountCoupons:     infoDomainCoupon.AmountCoupons,
		PriceCoupon:       infoDomainCoupon.PriceCoupon,
		Total:             infoDomainCoupon.Total,
		Codes:             couponMetaData,
		InfoPayment: mongodb_model.LinkPaymentInfo{
			OrderNumber: infoPayment.OrderNumber,
			Amount:      infoPayment.Amount,
			Status:      infoPayment.Status,
			RefNumber:   infoPayment.RefNumber,
			Type:        infoPayment.Type,
			IdTx:        infoPayment.IdTx,
			UUID:        infoPayment.UUID,
			Link:        infoPayment.Links.CheckoutLink,
			SuccessLink: infoPayment.Links.LinkSuccess,
			FailedLink:  infoPayment.Links.LinkFailed,
		},
		CreatedAt: time.Now().String(),
	}

	_, err := collection.InsertOne(context.Background(), modelCoupon)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, err
	}

	return modelCoupon, nil
}

func (m *MongoDB) SearchInfoPaymentStorage(collectionName string, idReference string) (*mongodb_model.CouponModel, error) {
	var codeCoupon mongodb_model.CouponModel
	collection := m.Database.Collection(collectionName)
	cleanIDReference := strings.TrimSpace(idReference)

	filter := bson.D{{"idDReferentProcess", cleanIDReference}}

	err := collection.FindOne(context.Background(), filter).Decode(&codeCoupon)
	if err != nil {
		m.log.ErrorLog.Print(err)
		return nil, err
	}

	return &codeCoupon, nil
}

func (m *MongoDB) UpdateStatusUalaStorage(collectionName string, idReference string, status string) error {
	collection := m.Database.Collection(collectionName)
	filter := bson.D{{"idDReferentProcess", idReference}}
	update := bson.D{
		{"$set", bson.D{
			{"infopayment.status", status},
		}},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		m.log.ErrorLog.Print(err)
		return err
	}

	return nil
}

func (m *MongoDB) SearchCodeStorage(collectionName string, codeName string) (*mongodbModel.CodeCoupon, error) {
	var codeCoupon mongodbModel.CodeCoupon

	collection := m.Database.Collection(collectionName)

	filter := bson.D{{"code", codeName}}

	err := collection.FindOne(context.Background(), filter).Decode(&codeCoupon)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return &codeCoupon, err
	}

	return &codeCoupon, nil
}

func (m *MongoDB) UpdateCouponStorage(collectionName string, couponUpdated *mongodbModel.CodeCoupon, codeName string) (*mongodb_model.CodeCoupon, error) {
	filter := bson.D{{"code", codeName}}
	collection := m.Database.Collection(collectionName)
	var couponModel mongodb_model.CodeCoupon
	updated := bson.D{
		{"$set", bson.D{
			{"isUsed", couponUpdated.IsUsed},
			{"cvu", couponUpdated.CVU},
			{"alias", couponUpdated.Alias},
			{"wallet", couponUpdated.Wallet},
			{"red", couponUpdated.Red},
		}},
	}
	_, err := collection.UpdateOne(context.Background(), filter, updated)

	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, err
	}

	err = collection.FindOne(context.Background(), filter).Decode(&couponModel)

	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, err
	}

	return &couponModel, nil
}

func (m *MongoDB) SearchCouponsAllUser(collectionName string, idReference string) (*mongodb_model.CouponModel, error) {
	var couponModel mongodb_model.CouponModel
	collection := m.Database.Collection(collectionName)
	filter := bson.D{{"idDReferentProcess", idReference}}
	err := collection.FindOne(context.Background(), filter).Decode(&couponModel)
	if err != nil {
		m.log.ErrorLog.Println(err)
		return nil, err
	}

	return &couponModel, nil
}
