import { extendTheme } from "@chakra-ui/react";

import { colors } from "./colors";

export const smallDevice = "(max-width: 768px)";

const theme = extendTheme({
  colors,
  colorSchemes: {},
  fonts: {
    heading: "Poppins",
    body: "Poppins",
  },
});

export default theme;
