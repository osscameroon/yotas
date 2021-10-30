import { useContext } from "react";
import { LangType, LocaleContext } from "./local-provider";
import STRINGS, { TranslateKeys } from "./strings";

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
