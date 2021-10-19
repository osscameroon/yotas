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

  const [lang, setLang] = useState(
    (langs.includes(navigatorLang) ? navigatorLang : "en") as LangType
  );

  const changeLang = (newLang: LangType) => {
    setLang(newLang);
  };

  return (
    <LocaleContext.Provider value={{ lang, changeLang }}>
      {children}
    </LocaleContext.Provider>
  );
};

export default LocaleProvider;
