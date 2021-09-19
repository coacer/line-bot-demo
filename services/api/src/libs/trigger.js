const { Proto } = require('../plugins/grpc');

module.exports = class Trigger extends Proto {
  constructor() {
    super();
    this.config = {
      HOST: process.env.TRIGGER_HOST,
      PORT: process.env.TRIGGER_PORT,
    };
    this.name = 'Trigger';
    this.loadClient();
  }

  test() {
    return this.invoke('test', {});
  }
};
