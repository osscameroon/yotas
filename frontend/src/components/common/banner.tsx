import {
  Box,
  Text,
  Button,
  useMediaQuery,
  Heading,
  SimpleGrid,
} from "@chakra-ui/react";
import React from "react";
import { AiFillGithub } from "react-icons/ai";
import BannerIllustration from "../../assets/components/banner-illustration";
import useTranslate from "../../locale/use-translate";

const Banner = () => {
  const { t } = useTranslate();

  const [isMd] = useMediaQuery("(min-width: 720px)");

  return (
    <SimpleGrid columns={isMd ? 2 : 1} spacing={10}>
      <Box
        alignItems={isMd ? "start" : "center"}
        display="flex"
        flexDirection="column"
        height="400px"
        justifyContent="space-evenly"
        textAlign={isMd ? "left" : "center"}
      >
        <Heading as="h1" color="primary" fontSize={isMd ? "5xl" : "3xl"}>
          {t("bannerMessage")}
        </Heading>
        <Text fontSize="xl">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Bibendum orci
          tellus phasellus donec eu aliquet aliquam ipsum feugiat eget orci.
        </Text>
        <Button
          bg="primary"
          color="white"
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
  );
};

export default Banner;
