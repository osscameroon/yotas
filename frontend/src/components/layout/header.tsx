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
  useMediaQuery,
  IconButton,
  Drawer,
  DrawerOverlay,
  DrawerContent,
  DrawerCloseButton,
  DrawerBody,
  DrawerHeader,
  useDisclosure,
} from "@chakra-ui/react";
import Logo from "../../assets/components/logo";
import { AiFillGithub, AiOutlineMenu } from "react-icons/ai";
import { BsTranslate } from "react-icons/bs";
import { LocaleContext } from "../../locale/local-provider";
import useTranslate from "../../locale/use-translate";
import { smallDevice } from "../../themes";
import SmallLogo from "../../assets/components/small-logo";

const Header = () => {
  const { t } = useTranslate();
  const [isSmallDevice] = useMediaQuery(smallDevice);

  const localeContext = useContext(LocaleContext);

  const handleEnglishSelected = () => localeContext.changeLang("en");
  const handleFrenchSelected = () => localeContext.changeLang("fr");

  const { isOpen, onClose, onOpen } = useDisclosure();

  //const menuButtonRef = useRef();

  return (
    <>
      {!isSmallDevice && (
        <Box boxShadow="sm">
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
              <Flex
                alignItems="center"
                flexDirection="row"
                justifyContent="flex-end"
              >
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
                      <MenuItem onClick={handleFrenchSelected}>
                        Français
                      </MenuItem>
                      <MenuItem onClick={handleEnglishSelected}>
                        English
                      </MenuItem>
                    </MenuList>
                  </Portal>
                </Menu>
              </Flex>
            </SimpleGrid>
          </Container>
        </Box>
      )}

      {isSmallDevice && (
        <Box boxShadow="sm" py="12px">
          <Container maxW="container.sm">
            <SimpleGrid columns={2} spacing={5}>
              <Box>
                <SmallLogo />
              </Box>
              <Box alignItems="center" display="flex" justifyContent="flex-end">
                <Menu>
                  <MenuButton color="primary.500">
                    <BsTranslate size="1.5em" />
                  </MenuButton>
                  <Portal>
                    <MenuList>
                      <MenuItem onClick={handleFrenchSelected}>
                        Français
                      </MenuItem>
                      <MenuItem onClick={handleEnglishSelected}>
                        English
                      </MenuItem>
                    </MenuList>
                  </Portal>
                </Menu>
                <Box mx="1" />
                <Box>
                  <IconButton
                    aria-label="menu button"
                    colorScheme="primary"
                    icon={<AiOutlineMenu />}
                    variant="outline"
                    onClick={onOpen}
                  />
                  <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
                    <DrawerOverlay />
                    <DrawerContent>
                      <DrawerCloseButton />
                      <DrawerHeader />
                      <DrawerBody>
                        <Flex
                          alignItems="center"
                          flexDirection="column"
                          fontWeight="bold"
                          height="200px"
                          justifyContent="space-between"
                        >
                          <Link href="#how">{t("howItWorks")}</Link>
                          <Link href="#about">{t("aboutUs")}</Link>
                          <Link href="#">{t("documentation")}</Link>
                          <Link href="#">{t("organisations")}</Link>
                        </Flex>
                      </DrawerBody>
                    </DrawerContent>
                  </Drawer>
                </Box>
              </Box>
            </SimpleGrid>
          </Container>
        </Box>
      )}
    </>
  );
};

export default Header;
