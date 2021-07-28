import * as React from "react";
import useTranslate from "../../locale/useTranslate";

const Header = () => {
  const { t } = useTranslate();
  return <> {t("home")} </>;
};

export default Header;
