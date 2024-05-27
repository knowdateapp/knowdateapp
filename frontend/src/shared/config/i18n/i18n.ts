import i18n, { ParseKeys } from 'i18next';
import resourcesToBackend from 'i18next-resources-to-backend';
import { initReactI18next } from 'react-i18next';

//TODO из-за плагина @conarti/eslint-plugin-feature-sliced падает еслинт в этом файле. Причина динамический импорт.

const defaultNS = 'translation';
type DefaultNS = typeof defaultNS;

export type TranslationKey = ParseKeys<DefaultNS, Record<string, string>, DefaultNS>;
const loadNameSpace = (language: string, namespace: string) =>
  /* eslint-disable next-line @conarti/feature-sliced/layers-slices */
  import(`./locales/${language}/${namespace}.json`);

void i18n
  .use(initReactI18next)
  .use(resourcesToBackend(loadNameSpace))
  .init({
    defaultNS,
    fallbackLng: 'en',
    debug: true,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
