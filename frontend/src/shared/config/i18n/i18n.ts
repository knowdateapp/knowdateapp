import i18n, { ParseKeys } from 'i18next';
import resourcesToBackend from 'i18next-resources-to-backend';
import { initReactI18next } from 'react-i18next';

const defaultNS = 'translation';
type DefaultNS = typeof defaultNS;

export type TranslationKey = ParseKeys<DefaultNS, Record<string, string>, DefaultNS>;

void i18n
  .use(initReactI18next)
  .use(
    resourcesToBackend(
      (language: string, namespace: string) => import(`./locales/${language}/${namespace}.json`),
    ),
  )
  .init({
    defaultNS,
    fallbackLng: 'en',
    debug: true,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
