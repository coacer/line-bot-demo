/**
 * エラーハンドラ
 */
module.exports = (err, req, res, next) => {
  console.warn('err: ', err);
  res.status(err.status || 500).json({
    code: err.code, // gRPC error code
    status: err.status || 500,
    message: err.message,
  });
};
