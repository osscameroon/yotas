import { frStrings } from "./fr";
import { enStrings } from "./en";

export type TranslateKeys = keyof typeof enStrings;

const STRINGS = {
  fr: frStrings,
  en: enStrings,
};

export default STRINGS;
