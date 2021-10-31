import React, { useContext } from "react";
import {
  Box,
  Container,
  Flex,
  SimpleGrid,
  Button,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Portal,
  Link,
} from "@chakra-ui/react";
import Logo from "../../assets/components/logo";
import { AiFillGithub } from "react-icons/ai";
import { BsTranslate } from "react-icons/bs";
import { LocaleContext } from "../../locale/local-provider";
import useTranslate from "../../locale/use-translate";

const Header = () => {
  const { t } = useTranslate();

  const localeContext = useContext(LocaleContext);

  const handleEnglishSelected = () => localeContext.changeLang("en");
  const handleFrenchSelected = () => localeContext.changeLang("fr");

  return (
    <Container maxW="container.xl" paddingY={5}>
      <SimpleGrid columns={2}>
        <Flex
          alignItems="center"
          flexDirection="row"
          fontWeight="bold"
          justifyContent="space-between"
        >
          <Logo />
          <Link href="#how">{t("howItWorks")}</Link>
          <Link href="#about">{t("aboutUs")}</Link>
          <Link href="#">{t("documentation")}</Link>
          <Link href="#">{t("organisations")}</Link>
        </Flex>
        <Flex alignItems="center" flexDirection="row" justifyContent="flex-end">
          <Button colorScheme="primary" rightIcon={<AiFillGithub />}>
            Sign Up
          </Button>
          <Box mx="1" />
          <Menu>
            <MenuButton color="primary.500">
              <BsTranslate size="2em" />
            </MenuButton>
            <Portal>
              <MenuList>
                <MenuItem onClick={handleFrenchSelected}>Français</MenuItem>
                <MenuItem onClick={handleEnglishSelected}>English</MenuItem>
              </MenuList>
            </Portal>
          </Menu>
        </Flex>
      </SimpleGrid>
    </Container>
  );
};

export default Header;
