import React, { useState } from "react";

type LocaleContextType = {
  lang: string;
  changeLang: (newLang: string) => void;
};

const DefaultLang: LocaleContextType = {
  lang: "",
  changeLang: () => null,
};

export const LocaleContext = React.createContext(DefaultLang);

const LocaleProvider = ({ children }: { children: JSX.Element }) => {
  const [lang, setLang] = useState(window.navigator.language.split("-")[0]);
  const changeLang = (newLang: string) => {
    setLang(newLang);
  };

  return (
    <LocaleContext.Provider value={{ lang, changeLang }}>
      {children}
    </LocaleContext.Provider>
  );
};

export default LocaleProvider;
