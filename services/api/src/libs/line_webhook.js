const { Proto } = require('../plugins/grpc');

module.exports = class LineWebhook extends Proto {
  constructor() {
    super();
    this.config = {
      HOST: process.env.WEBHOOK_HOST,
      PORT: process.env.WEBHOOK_PORT,
    };
    this.name = 'LineWebhook';
    this.loadClient();
  }

  health() {
    return this.invoke('health', {});
  }

  /**
   * メッセージイベントハンドラ
   * @param {*} event
   */
  message(event) {
    return this.invoke('message', { event });
  }
};
