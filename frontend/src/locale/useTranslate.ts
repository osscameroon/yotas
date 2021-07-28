import { useContext } from "react";
import { LocaleContext } from "./LocaleProvider";
import STRINGS from "./strings";

function translate(lang: string, keyword: string) {
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
