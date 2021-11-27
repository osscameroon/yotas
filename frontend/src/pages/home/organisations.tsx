import React from "react";
import {
  Box,
  Heading,
  Button,
  SimpleGrid,
  Text,
  Input,
  Container,
  InputGroup,
  InputRightElement,
} from "@chakra-ui/react";
import { SearchIcon } from "@chakra-ui/icons";
import OrganisationCard from "../../components/common/organisation-card";
import { ORGANISATIONS } from "../../fixtures/organisations";
import useTranslate from "../../locale/use-translate";

const Organisations = () => {
  const { t } = useTranslate();

  return (
    <Box py="3em">
      <Container maxW="container.xl">
        <Heading textAlign="center">{t("joinOrganization")}</Heading>
        <Box m="2em" />
        <Text textAlign="center">{t("joinGoal")}</Text>
        <Box m="2em" />
        <Box display="flex" justifyContent="center">
          <InputGroup width="500px">
            <Input placeholder={`${t("orgSearch")}`} variant="outline" />
            <InputRightElement>
              <SearchIcon cursor="pointer" />
            </InputRightElement>
          </InputGroup>
        </Box>
        <Box m="2em" />
        <SimpleGrid columns={{ sm: 1, md: 3 }} space={20}>
          {ORGANISATIONS.map((org, i) => (
            <OrganisationCard key={i} {...org} />
          ))}
        </SimpleGrid>
        <Box m="2em" />
        <Box textAlign="center">
          <Button colorScheme="primary">{t("viewMore")}...</Button>
        </Box>
      </Container>
    </Box>
  );
};

export default Organisations;
