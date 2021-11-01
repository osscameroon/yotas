import React from "react";
import { Box, Flex, Image, Text } from "@chakra-ui/react";

type ArticleCardProps = {
  picture: string;
  name: string;
  price: number;
};

const ArticleCard = ({ name, picture, price }: ArticleCardProps) => {
  return (
    <Box borderRadius="lg" boxShadow="sm" cursor="pointer">
      <Box>
        <Image htmlHeight="300px" htmlWidth="auto" src={picture} />
      </Box>
      <Flex alignItems="center" justifyContent="space-between" padding="15px">
        <Text fontSize="xl" fontWeight="bold">
          {name}
        </Text>
        <Text color="primary.500" fontSize="sm" fontWeight="bold">
          {price} Y
        </Text>
      </Flex>
    </Box>
  );
};

export default ArticleCard;
