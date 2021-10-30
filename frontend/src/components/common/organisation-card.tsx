import React from "react";
import {
  Box,
  Heading,
  Text,
  Button,
  Flex,
  Image,
  Link,
} from "@chakra-ui/react";

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
  return (
    <Box boxShadow="lg" height="420px" margin="20px" rounded="lg">
      <Box borderRadius="2xl" height="150px">
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
        mt="2em"
        padding="15px"
      >
        <Box>
          <Heading fontSize="xl" fontWeight="bold">
            {name}
          </Heading>
          <Box mt="3" />
          <Text>{description}</Text>
        </Box>
        <Flex alignItems="center" justifyContent="space-between">
          <Text>{membersCount} Members</Text>
          <Link href={`#${id}`}>
            <Button colorScheme="primary"> Join </Button>
          </Link>
        </Flex>
      </Flex>
    </Box>
  );
};

export default OrganisationCard;
