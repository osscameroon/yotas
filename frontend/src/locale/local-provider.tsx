import React, { useState } from "react";

export type LangType = "fr" | "en";
const langs = ["fr", "en"];

type LocaleContextType = {
  lang: LangType;
  changeLang: (newLang: LangType) => void;
};

const DefaultLang: LocaleContextType = {
  lang: "fr",
  changeLang: () => null,
};

export const LocaleContext = React.createContext(DefaultLang);

const LocaleProvider = ({ children }: { children: JSX.Element }) => {
  const [navigatorLang] = window.navigator.language.split("-");
  const langKey = "defaultLanguage";
  const savedLang = localStorage.getItem(langKey);

  const languageToUse = (): LangType => {
    if (savedLang) return savedLang as LangType;
    if (langs.includes(navigatorLang)) return navigatorLang as LangType;

    return "en";
  };

  const [lang, setLang] = useState(languageToUse());

  const changeLang = (newLang: LangType) => {
    localStorage.setItem(langKey, newLang);
    setLang(newLang);
  };

  return (
    <LocaleContext.Provider value={{ lang, changeLang }}>
      {children}
    </LocaleContext.Provider>
  );
};

export default LocaleProvider;
