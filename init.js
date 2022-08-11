// admin
db.auth("root", "example");

// user
userdb = db.getSiblingDB("UserService");
userdb.createUser({
  user: "ks",
  pwd: "ks",
  roles: [{ role: "readWrite", db: "UserService" }],
  mechanisms: ["SCRAM-SHA-1"],
  passwordDigestor: "client",
});
userdb.auth("ks", "ks");
