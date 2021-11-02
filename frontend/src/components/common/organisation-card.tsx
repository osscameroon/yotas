import React from "react";
import {
  Box,
  Heading,
  Text,
  Button,
  Flex,
  Image,
  Link,
  useMediaQuery,
} from "@chakra-ui/react";
import useTranslate from "../../locale/use-translate";
import { smallDevice } from "../../themes";

type OrganisationCardProps = {
  name: string;
  description: string;
  membersCount: number;
  id: string;
  logo: string;
  banner: string;
};

const OrganisationCard = ({
  banner,
  description,
  id,
  logo,
  membersCount,
  name,
}: OrganisationCardProps) => {
  const { t } = useTranslate();
  const [isSmallDevice] = useMediaQuery(smallDevice);

  return (
    <Box
      boxShadow="lg"
      height={isSmallDevice ? "400" : "420px"}
      margin="20px"
      rounded="lg"
    >
      <Box height="150px">
        <Box m={3} position="absolute">
          <Image alt={`logo-${name}`} boxSize="5em" src={logo} />
        </Box>
        <Box height="200px">
          <Image
            alt={`banner-${name}`}
            htmlHeight="auto"
            htmlWidth="100%"
            src={banner}
          />
        </Box>
      </Box>
      <Flex
        flexDirection="column"
        height="250px"
        justifyContent="space-around"
        mt={isSmallDevice ? "0em" : "2em"}
        padding="15px"
      >
        <Box>
          <Heading fontSize="xl" fontWeight="bold">
            {name}
          </Heading>
          <Box mt="1em" />
          <Text>{description}</Text>
        </Box>
        <Flex alignItems="center" justifyContent="space-between">
          <Text>
            {membersCount} {t("members")}
          </Text>
          <Link href={`#${id}`}>
            <Button colorScheme="primary">{t("join")}</Button>
          </Link>
        </Flex>
      </Flex>
    </Box>
  );
};

export default OrganisationCard;
