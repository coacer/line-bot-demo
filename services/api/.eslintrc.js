module.exports = {
  env: {
    node: true,
    es6: true,
  },
  extends: ['eslint:recommended', 'plugin:node/recommended', 'plugin:prettier/recommended'],
  rules: {
    'no-console': 'off',
    'no-unused-vars': 'off',
    'no-empty': 'warn',
  },
};
