import * as React from "react";
import { Box } from "@chakra-ui/react";

import Footer from "./footer";
import Header from "./header";

type LayoutProps = {
  children: React.ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
  return (
    <>
      <Header />
      <Box marginY={5} />
      <main>{children}</main>
      <Footer />
    </>
  );
};

export default Layout;
