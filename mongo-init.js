db.createUser(
    {
        user: process.env.MONGO_INITDB_ROOT_USERNAME,
        pwd: process.env.MONGO_INITDB_ROOT_PASSWORD,
        roles: [
            {
                role: "readWrite",
                db: process.env.MONGO_INITDB_DATABASE
            }
        ]
    }
);
db.createCollection("users");
db.getCollection("users").createIndex( { "username": 1 }, { unique: true } );
db.getCollection("users").createIndex( { "email": 1 }, { unique: true } );