## Ejecutar Go cmd/app
`MongoDB`
> go run . -enginedb mongodb -portdb 27017 -userdb almendra -passwordb 1ASWWWaeq -dbname microservicio
`MySQL`
>go run . -enginedb mysql -portdb 33060 -userdb root -passwordb secret
## Ejecutar contenedor docker
`MongoDB`
> sudo docker run -d -p 27017:27017 --name mongodb-container -e MONGO_INITDB_ROOT_USERNAME=almendra -e MONGO_INITDB_ROOT_PASSWORD="1ASWWWaeq" -v mongodb-data:/data/db mongo
`MySQL`
>sudo docker run -d -p 33060:3306 --name mysql-db  -e MYSQL_ROOT_PASSWORD=secret --mount src=mysql-db-data,dst=/var/lib/mysql mysql
>docker run -d -p 33060:3306 --name mysql-db  -e MYSQL_ROOT_PASSWORD=secret --mount src=mysql-db-data,dst=/var/lib/mysql --network microservice_service-mongodb mysql 

### En caso de este error
>docker: Error response from daemon: Conflict. The container name "/mongodb-container" is already in use by container "c2d38b2113b05c511f85526835a2b02de0604623f66cb567976db28f4a6d5496". You have to remove (or rename) that container to be able to reuse that name.

### Hacer los siguientes pasos`
`docker stop mongodb-container`
`docker rm mongodb-container`
`Ejecutar otra vez el comando docker run ...`

### Recurso de Arquitectura Hexagonal
>https://carlos.lat/blog/hexagonal-architecture-using-go-fiber/


### Comandos Docker
>docker compose build --no-cache
>docker compose up -d
>docker logs Nombre_Imagen

`Es necesario crerar la DB en Workbench`
>CREATE DATABASE users;

### Ejemplo de como se guarda el BODY 
> localhost:4000/v1/giftcard/create
```
[
  {
    _id: ObjectId('666cd93a2f6d17734f7dae22'),
    idOwner: '666495d87266690cbddd165c',
    title: 'cel',
    description: 'Esto es una descripcion',
    amountCoupons: '0',
    priceCoupon: '100',
    total: '100',
    codes: [
      {
        code: 'Eu9aqPBJ',
        expireat: ISODate('2024-07-14T23:58:50.961Z'),
        beneficiaryuser: {
          userBeneficiaryID: ObjectId('000000000000000000000000'),
          name: '',
          code: 0,
          mobile: 0,
          email: ''
        },
        isused: false,
        price: 100,
        cvu: '',
        alias: '',
        wallet: '',
        red: ''
      }
    ],
    infopayment: {
      link: 'https://uala-checkout.preprod.geopagos.com/orders/6ff3ce7c-1bdf-4256-8e11-9f767d5452ab',
      successlink: 'https://www.utl-test.com/search?q=failed',
      failedlink: '',
      ordernumber: '0004216-0000230144',
      amount: '100',
      refnumber: 'b6ab3e6f-686f-4f03-873b-11f90cbb6daa',
      status: 'PENDING',
      type: 'Order',
      idtx: '/api/v2/orders/6ff3ce7c-1bdf-4256-8e11-9f767d5452ab',
      uuid: '6ff3ce7c-1bdf-4256-8e11-9f767d5452ab'
    }
  }
]
```