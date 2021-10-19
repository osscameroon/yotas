import * as React from "react";
import { Box } from "@chakra-ui/react";

import Banner from "../components/common/banner";
import Layout from "../components/layout/layout";

const Home = () => {
  return (
    <Layout>
      <Box>
        <Banner />
      </Box>
    </Layout>
  );
};

export default Home;
