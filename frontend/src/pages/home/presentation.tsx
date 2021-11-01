import React from "react";
import {
  Box,
  Heading,
  Text,
  Flex,
  Container,
  Icon,
  SimpleGrid,
} from "@chakra-ui/react";
import { YotasIllustration } from "../../assets/components/yotas-illustration";
import { BsDoorOpen, BsAlignStart, BsShopWindow } from "react-icons/bs";
import { FaCoins } from "react-icons/fa";
import { IconType } from "react-icons";
import useTranslate from "../../locale/use-translate";

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
          <Box my="2em" />
          <SimpleGrid columns={2} spacing={20}>
            <Box>
              <Text ml={10}>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Bibendum orci tellus phasellus donec eu aliquet aliquam ipsum
                feugiat eget orci.Lorem ipsum dolor sit amet, consectetur
                adipiscing elit. Bibendum orci tellus phasellus donec eu aliquet
                aliquam ipsum feugiat eget orci.
              </Text>
              <Box my="4" />
              <YotasIllustration />
            </Box>
            <Flex flexDirection="column" justifyContent="space-around">
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
