import React from "react";
import {
  Box,
  Text,
  Button,
  useMediaQuery,
  Heading,
  SimpleGrid,
  Container,
} from "@chakra-ui/react";
import { AiFillGithub } from "react-icons/ai";
import BannerIllustration from "../../assets/components/banner-illustration";
import useTranslate from "../../locale/use-translate";

const Banner = () => {
  const { t } = useTranslate();

  const [isMd] = useMediaQuery("(min-width: 768px)");

  return (
    <Container maxW="container.xl">
      <SimpleGrid columns={isMd ? 2 : 1} spacing={10}>
        <Box
          alignItems={isMd ? "start" : "center"}
          display="flex"
          flexDirection="column"
          height="400px"
          justifyContent="space-evenly"
          textAlign={isMd ? "left" : "center"}
        >
          <Heading as="h1" color="primary.500" fontSize={isMd ? "5xl" : "3xl"}>
            {t("bannerMessage")}
          </Heading>
          <Text fontSize="xl">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Bibendum
            orci tellus phasellus donec eu aliquet aliquam ipsum feugiat eget
            orci.
          </Text>
          <Button
            colorScheme="primary"
            rightIcon={<AiFillGithub />}
            variant="solid"
          >
            {t("joinButtonText")}
          </Button>
        </Box>
        {isMd && (
          <Box display="flex" justifyContent="center">
            <BannerIllustration />
          </Box>
        )}
      </SimpleGrid>
    </Container>
  );
};

export default Banner;
