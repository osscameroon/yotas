import * as React from "react";
import { Container } from "@chakra-ui/react";

import Footer from "./footer";
import Header from "./header";

type LayoutProps = {
  children: React.ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
  return (
    <>
      <Header />

      <Container maxW="container.xl">
        <main>{children}</main>
      </Container>

      <Footer />
    </>
  );
};

export default Layout;
