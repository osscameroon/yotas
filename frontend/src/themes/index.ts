import { extendTheme } from "@chakra-ui/react";

import { colors } from "./colors";

const theme = extendTheme({
  colors,
  colorSchemes: {},
  fonts: {
    heading: "Poppins",
    body: "Poppins",
  },
});

export default theme;
