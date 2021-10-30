import * as React from "react";
import { Box } from "@chakra-ui/react";

import Banner from "./banner";
import Layout from "../../components/layout/layout";
import Presentation from "./presentation";
import Organisations from "./organisations";

const Home = () => {
  return (
    <Layout>
      <Box>
        <Banner />
        <Box my={2} />
        <Presentation />
        <Box my={2} />
        <Organisations />
      </Box>
    </Layout>
  );
};

export default Home;
