const path = require('path');
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

class ProtoLoader {
  constructor(
    name,
    options = {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true,
    }
  ) {
    this.name = name;
    this.lowercaseName = name.toLowerCase();
    this.options = options;
  }

  getClient() {
    const PROTO_PATH = path.resolve(__dirname, `../../../../rpc/${this.lowercaseName}.proto`);
    const packageDefinition = protoLoader.loadSync(PROTO_PATH, this.options);
    const proto = grpc.loadPackageDefinition(packageDefinition);

    return new proto[this.lowercaseName][this.name]('127.0.0.1:5000', grpc.credentials.createInsecure());
  }
}

exports.ProtoLoader = ProtoLoader;
