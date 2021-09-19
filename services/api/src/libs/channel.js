const { Proto } = require('../plugins/grpc');

module.exports = class Channel extends Proto {
  constructor() {
    super();
    this.config = {
      HOST: process.env.CHANNEL_HOST,
      PORT: process.env.CHANNEL_PORT,
    };
    this.name = 'Channel';
    this.loadClient();
  }

  test() {
    return this.invoke('test', {});
  }
};
