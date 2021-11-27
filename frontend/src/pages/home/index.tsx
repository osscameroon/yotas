import * as React from "react";
import {
  Box,
  SimpleGrid,
  Container,
  Image,
  Heading,
  Text,
  Button,
  Flex,
  useMediaQuery,
} from "@chakra-ui/react";

import Banner from "./banner";
import Layout from "../../components/layout/layout";
import Presentation from "./presentation";
import Organisations from "./organisations";
import useTranslate from "../../locale/use-translate";

import docIllustration from "../../assets/images/docs.png";
import logoBackground from "../../assets/images/logo-background.png";
import ARTICLES from "../../fixtures/home-articles";
import ArticleCard from "../../components/common/article-card";
import { smallDevice } from "../../themes";

const Home = () => {
  const { t } = useTranslate();
  const [isSmallDevice] = useMediaQuery(smallDevice);

  document.title = t("homePage");

  return (
    <Layout>
      <Box>
        <Banner />
        <Box my={2} />
        <Presentation />
        <Box my={2} />
        <Organisations />
        <Box my={2} />
        <Box
          bgColor="lighterGrey"
          bgImage={logoBackground}
          bgPos="right"
          bgRepeat="no-repeat"
        >
          <Container maxW="container.xl" py="3em">
            <SimpleGrid columns={{ sm: 1, md: 2 }}>
              <Box
                alignItems="center"
                display={isSmallDevice ? "none" : "flex"}
                justifyContent="center"
              >
                <Image boxSize="470px" src={docIllustration} />
              </Box>
              <Box
                display="flex"
                flexDirection="column"
                justifyContent="center"
              >
                <Box>
                  <Heading>{t("documentation")}</Heading>
                  <Box mt={5} />
                  <Text>{t("docText")}</Text>
                  <Box mt={5} />
                  <Button colorScheme="primary">{t("viewDoc")}</Button>
                </Box>
              </Box>
            </SimpleGrid>
          </Container>
        </Box>
        <Box my={2} />
        <Box py="3em">
          <Container
            alignItems="center"
            display="flex"
            flexDirection="column"
            maxW="container.xl"
          >
            <Heading>{t("shopArticle")}</Heading>
            <Box mt="3em" />
            <SimpleGrid columns={{ sm: 1, md: 3 }} spacing={10}>
              {ARTICLES.map((article, i) => (
                <ArticleCard {...article} key={i} />
              ))}
            </SimpleGrid>
            <Box mt="3em" />
            <Button colorScheme="primary">{t("visitShop")}</Button>
          </Container>
        </Box>
        <Box my={2} />
        <Box bgColor="lighterGrey" py="3em">
          <Container maxW="container.xl">
            <SimpleGrid columns={{ sm: 1, md: 2 }} spacing={5}>
              <Flex
                alignItems="flex-start"
                direction="column"
                justifyContent="center"
              >
                <Text fontSize="xl" fontWeight="bold">
                  {t("donate")}
                </Text>
                <Box mt="1em" />
                <Text>{t("donateText")}</Text>
              </Flex>
              <Flex alignItems="center" justifyContent="flex-end">
                <Button colorScheme="primary">{t("donate")}</Button>
              </Flex>
            </SimpleGrid>
          </Container>
        </Box>
      </Box>
    </Layout>
  );
};

export default Home;
