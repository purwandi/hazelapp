const fs = require('fs');
const path = require('path');

const prettierOptions = JSON.parse(
  fs.readFileSync(path.resolve(__dirname, '.prettierrc'), 'utf8'),
);

module.exports = {
  parser: 'babel-eslint',
  env: {
    browser: true,
    es6: true
  },
  extends: ['airbnb', 'prettier', 'prettier/react'],
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly'
  },
  parserOptions: {
    ecmaFeatures: {
      jsx: true
    },
    ecmaVersion: 2018,
    sourceType: 'module'
  },
  plugins: ['react', 'prettier'],
  rules: {
    'arrow-body-style': [2, 'as-needed'],
    'consistent-return': 'off',
    'import/imports-first': 0,
    'import/newline-after-import': 0,
    'import/prefer-default-export': 0,
    'jsx-a11y/label-has-for': 0,
    'jsx-a11y/label-has-associated-control': [2, { controlComponents: ['Input'] }],
    'jsx-a11y/no-static-element-interactions': 0,
    'jsx-a11y/click-events-have-key-events': 0,
    'max-len': ['error', { code: 120 }],
    'prettier/prettier': ['error', prettierOptions],
    'react/forbid-prop-types': 0,
    'react/jsx-filename-extension': [1, { extensions: ['.js', '.jsx'] }],
    'react/jsx-props-no-spreading': 0,
    'react/jsx-uses-vars': 2,
    'react/prop-types': ['error', {
      ignore: ['navigation', 'match', 'history']
    }]
  }
};
