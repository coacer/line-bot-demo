const path = require('path');
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

class Proto {
  constructor(
    options = {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true,
    }
  ) {
    this.name = null;
    this.config = { HOST: null, PORT: null };
    this.options = options;
    this.client = null;
  }

  loadClient() {
    if (!this.name) return;
    const lowercaseName = this.name.toLowerCase();
    const PROTO_PATH = path.resolve(__dirname, `../../rpc/${lowercaseName}.proto`);
    const packageDefinition = protoLoader.loadSync(PROTO_PATH, this.options);
    this.proto = grpc.loadPackageDefinition(packageDefinition);
    this.client = new this.proto[lowercaseName][this.name](
      `${this.config.HOST}:${this.config.PORT}`,
      grpc.credentials.createInsecure()
    );
  }

  invoke(fn, args) {
    return new Promise((resolve, reject) => {
      try {
        if (!this.client) throw new Error('Not load client.');
        this.client[fn](args, (err, res) => {
          console.log(`Invoke ${this.name}.${fn}`);
          if (err) {
            reject(err);
          } else {
            resolve(res);
          }
        });
      } catch (e) {
        reject(e);
      }
    });
  }
}

exports.Proto = Proto;
