import React from "react";
import {
  Box,
  Heading,
  Text,
  Flex,
  Container,
  Icon,
  SimpleGrid,
  useMediaQuery,
} from "@chakra-ui/react";
import { YotasIllustration } from "../../assets/components/yotas-illustration";
import { BsDoorOpen, BsAlignStart, BsShopWindow } from "react-icons/bs";
import { FaCoins } from "react-icons/fa";
import { IconType } from "react-icons";
import useTranslate from "../../locale/use-translate";
import { smallDevice } from "../../themes";

type IconTextProps = {
  icon: IconType;
  text: string;
  boxSize: string;
  color: string;
};

const IconText = ({ boxSize, color, icon, text }: IconTextProps) => (
  <Flex alignItems="center" justifyContent="start">
    <Icon as={icon} boxSize={boxSize} color={color} /> <Box mx="2" />{" "}
    <Text fontWeight="bold">{text} </Text>
  </Flex>
);

const Presentation = () => {
  const { t } = useTranslate();
  const [isSmallDevice] = useMediaQuery(smallDevice);

  const textMargin = isSmallDevice ? 0 : 20;

  const iconStyle = {
    boxSize: "45px",
    color: "primary.500",
  };

  return (
    <Box bg="lighterGrey" py="3em" width="100%">
      <Container maxW="container.xl">
        <Flex flexDirection="column" id="#how">
          <Heading as="h1" textAlign="center">
            {t("howDoesWork")}
          </Heading>
          <Box my={isSmallDevice ? "1em" : "2em"} />
          <SimpleGrid
            columns={{ sm: 1, md: 2 }}
            spacing={isSmallDevice ? 5 : 20}
          >
            <Box>
              <Text ml={textMargin}>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Bibendum orci tellus phasellus donec eu aliquet aliquam ipsum
                feugiat eget orci.Lorem ipsum dolor sit amet, consectetur
                adipiscing elit. Bibendum orci tellus phasellus donec eu aliquet
                aliquam ipsum feugiat eget orci.
              </Text>
              <Box my={isSmallDevice ? "1em" : "2em"} />
              <YotasIllustration />
            </Box>
            <Flex
              alignItems="flex-start"
              flexDirection="column"
              height={isSmallDevice ? "250px" : "auto"}
              justifyContent="space-around"
              ml={isSmallDevice ? "3em" : "0em"}
            >
              <IconText
                icon={BsDoorOpen}
                {...iconStyle}
                text={t("joinOrganization")}
              />
              <IconText
                icon={BsAlignStart}
                {...iconStyle}
                text={t("startContrib")}
              />
              <IconText icon={FaCoins} {...iconStyle} text={t("earnYotas")} />
              <IconText
                icon={BsShopWindow}
                {...iconStyle}
                text={t("shopOnYotas")}
              />
            </Flex>
          </SimpleGrid>
        </Flex>
      </Container>
    </Box>
  );
};

export default Presentation;
