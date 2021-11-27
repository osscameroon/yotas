import { useContext } from "react";
import { LangType, LocaleContext } from "./local-provider";
import { frStrings } from "./fr";
import { enStrings } from "./en";

export type TranslateKeys = keyof typeof enStrings;

const STRINGS = {
  fr: frStrings,
  en: enStrings,
};

function translate(lang: LangType, keyword: TranslateKeys): string {
  return STRINGS[lang][keyword];
}

function useTranslate() {
  const localeContext = useContext(LocaleContext);

  const t = (keyword: TranslateKeys) => {
    return translate(localeContext.lang, keyword);
  };

  return { t };
}

export default useTranslate;
