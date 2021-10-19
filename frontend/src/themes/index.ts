import { extendTheme } from "@chakra-ui/react";
import { createBreakpoints } from "@chakra-ui/theme-tools";

const breakpoints = createBreakpoints({
  sm: "320px",
  md: "768px",
  lg: "960px",
  xl: "1200px",
});

const theme = extendTheme({
  breakpoints,
  colors: {
    primary: "#488FF9",
    secondary: "#D6B1FF",
    accentGreen: "#099876",
    accentPink: "#E1118E",
  },
  colorSchemes: {},
  fonts: {
    heading: "Poppins",
    body: "Poppins",
  },
});

export default theme;
