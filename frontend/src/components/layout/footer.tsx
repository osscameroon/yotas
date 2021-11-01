import * as React from "react";
import {
  Box,
  Container,
  Flex,
  SimpleGrid,
  List,
  ListItem,
  Text,
  Link,
} from "@chakra-ui/react";
import useTranslate from "../../locale/use-translate";

const Footer = () => {
  const { t } = useTranslate();

  return (
    <Box bg="primary.500" color="white" py="3em">
      <Container maxW="container.xl">
        <SimpleGrid columns={4} spacing={10}>
          <Flex alignItems="center" flexDirection="column">
            <List spacing={3}>
              <ListItem>
                <Text fontWeight="bold">{t("links")}</Text>
              </ListItem>
              <ListItem>
                <Link href="https://github.com/osscameroon/yotas">
                  {t("githubRepo")}
                </Link>
              </ListItem>
              <ListItem>
                <Link href="#">{t("documentation")} </Link>
              </ListItem>
            </List>
          </Flex>
          <Flex alignItems="center" flexDirection="column">
            <List spacing={3}>
              <ListItem>
                <Text fontWeight="bold">{t("support")}</Text>
              </ListItem>
              <ListItem>
                <Link href="#">{t("howItWorks")}</Link>
              </ListItem>
              <ListItem>
                <Link href="#">{t("contact")}</Link>
              </ListItem>
            </List>
          </Flex>
          <Flex alignItems="center" flexDirection="column">
            <List spacing={3}>
              <ListItem>
                <Text fontWeight="bold">{t("socialMedias")}</Text>
              </ListItem>
              <ListItem>
                <Link href="#">Twitter</Link>
              </ListItem>
              <ListItem>
                <Link href="#">Telegram</Link>
              </ListItem>
              <ListItem>
                <Link href="#">Youtube</Link>
              </ListItem>
            </List>
          </Flex>
          <Flex alignItems="center" flexDirection="column">
            <List spacing={3}>
              <ListItem>
                <Text fontWeight="bold">{t("navigate")}</Text>
              </ListItem>
              <ListItem>
                <Link href="#">{t("home")}</Link>
              </ListItem>
              <ListItem>
                <Link href="#">{t("organisations")}</Link>
              </ListItem>
              <ListItem>
                <Link href="#">{t("donate")}</Link>
              </ListItem>
            </List>
          </Flex>
        </SimpleGrid>
        <Box mt="3em" textAlign="center">
          {t("madeBy")} <Link href="https://osscameroon.com">OSS Cameroon</Link>
        </Box>
      </Container>
    </Box>
  );
};

export default Footer;
