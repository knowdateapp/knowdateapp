module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended-type-checked',
    'plugin:@typescript-eslint/stylistic-type-checked',
    'plugin:react/recommended',
    'plugin:react-hooks/recommended',
    'prettier',
    'plugin:prettier/recommended',
    'plugin:import/recommended',
    'plugin:jsx-a11y/recommended',
    'plugin:promise/recommended',
    'plugin:@conarti/feature-sliced/recommended',
    'plugin:@tanstack/eslint-plugin-query/recommended',
    'plugin:optimize-regex/recommended',
    'plugin:react/jsx-runtime',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs', 'index.html', 'src/shared/config/i18n/i18n.ts'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: ['./tsconfig.json', './tsconfig.node.json'],
    tsconfigRootDir: __dirname,
  },
  plugins: ['react-refresh', 'import'],
  settings: {
    react: {
      version: 'detect',
    },
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts', '.tsx'],
    },
    'import/resolver': {
      typescript: {
        alwaysTryTypes: true,
        project: './tsconfig.json',
      },
    },
    rules: {
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
      'sort-imports': [
        'error',
        {
          ignoreCase: false,
          ignoreDeclarationSort: false,
          ignoreMemberSort: false,
          memberSyntaxSortOrder: ['none', 'all', 'multiple', 'single'],
          allowSeparatedGroups: false,
        },
      ],
      'react/jsx-uses-react': 'off',
      'import/no-unresolved': 'error',
      'no-console': 'error',
      '@conarti/feature-sliced/layers-slices': 'off',
    },
  },
};
