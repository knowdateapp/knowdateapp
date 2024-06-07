import 'i18next';
import translation from '../config/i18n/locales/ru/translation.json';

declare module 'i18next' {
  interface CustomTypeOptions {
    resources: {
      translation: typeof translation;
    };
  }
}
