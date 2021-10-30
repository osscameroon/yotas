import React from "react";
import {
  Box,
  Heading,
  Button,
  SimpleGrid,
  Text,
  Input,
  Container,
} from "@chakra-ui/react";
import OrganisationCard from "../../components/common/organisation-card";
import { ORGANISATIONS } from "../../fixtures/organisations";

const Organisations = () => {
  return (
    <Box py="3em">
      <Container maxW="container.xl">
        <Heading textAlign="center">Join Organisation</Heading>
        <Box m="2em" />
        <Text textAlign="center">
          Join organisation to contribute and start earning Yotas
        </Text>
        <Box m="2em" />
        <Box textAlign="center">
          <Input
            placeholder="Search for organisations"
            variant="outline"
            width="500px"
          />
        </Box>
        <Box m="2em" />
        <SimpleGrid columns={3} space={20}>
          {ORGANISATIONS.map((org, i) => (
            <OrganisationCard key={i} {...org} />
          ))}
        </SimpleGrid>
        <Box m="2em" />
        <Box textAlign="center">
          <Button colorScheme="primary">View more...</Button>
        </Box>
      </Container>
    </Box>
  );
};

export default Organisations;
