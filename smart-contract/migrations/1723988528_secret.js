const SecretStorage = artifacts.require("SecretStorage");

module.exports = function (deployer) {
  deployer.deploy(SecretStorage);
};
