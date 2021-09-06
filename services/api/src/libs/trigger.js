const { Proto } = require('../plugins/grpc');

module.exports = class Trigger extends Proto {
  constructor() {
    super();
    this.name = 'Trigger';
    this.loadClient();
  }

  test() {
    return this.invoke('test', {});
  }
};
