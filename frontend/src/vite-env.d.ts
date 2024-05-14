/// <reference types="vite/client" />

import 'i18next';
import translation from 'shared/config/i18n/locales/en/translation.json';

declare module 'i18next' {
  interface CustomTypeOptions {
    resources: {
      translation: typeof translation;
    };
  }
}
