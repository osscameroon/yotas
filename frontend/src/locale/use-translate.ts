import { useContext } from "react";
import { LangType, LocaleContext } from "./local-provider";
import STRINGS from "./strings";

function translate(lang: LangType, keyword: string): string {
  const keys: string[] = Object.keys(STRINGS);

  if (keys.includes(keyword)) return STRINGS[keyword][lang];
  return keyword;
}

function useTranslate() {
  const localeContext = useContext(LocaleContext);

  const t = (keyword: string) => {
    return translate(localeContext.lang, keyword);
  };

  return { t };
}

export default useTranslate;
