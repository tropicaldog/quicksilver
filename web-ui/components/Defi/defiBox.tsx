import React, { useState } from 'react';
import {
  Box,
  Button,
  Flex,
  Heading,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Text,
  Select,
  Stack,
  useColorModeValue,
  ButtonGroup,
  HStack,
} from '@chakra-ui/react';
import { ChevronDownIcon, ExternalLinkIcon } from '@chakra-ui/icons';
import { useDefiData } from '@/hooks/useQueries';
type ActionButtonTitle = 'Add Liquidity' | 'Borrow' | 'Lend' | 'Mint Stablecoin' | 'Vaults';
interface DefiAsset {
  id: string;
  assetPair: string;
  apy: number;
  tvl: string;
  provider: string;
  action: string;
}

const actionTitles: Record<string, ActionButtonTitle> = {
  'add-liquidity': 'Add Liquidity',
  borrow: 'Borrow',
  lend: 'Lend',
  'mint-stablecoin': 'Mint Stablecoin',
  vaults: 'Vaults',
};

interface DefiData {
  assetPair: string;
  apy: number;
  tvl: number;
  provider: string;
  action: string;
}

const filterCategories: Record<string, (data: DefiData) => boolean> = {
  All: () => true,
  'Borrowing & Lending': (data: DefiData) => data.action === 'Borrow' || data.action === 'Lend',
  Vaults: (data: DefiData) => data.action === 'Vaults',
  'Liquidity Providers': (data: DefiData) => data.action === 'Add Liquidity',
  'Mint Stable Coins': (data: DefiData) => data.action === 'Mint Stablecoin',
};

const formatApy = (apy: number) => {
  return `${(apy * 100).toFixed(2)}%`; // Converts to percentage and formats to 2 decimal places
};

const DefiTable = () => {
  const { defi, isLoading, isError } = useDefiData();

  const [activeFilter, setActiveFilter] = useState<string>('All');

  const handleFilterClick = (filter: string) => {
    setActiveFilter(filter);
  };

  const filteredData = defi ? defi.filter(filterCategories[activeFilter]) : [];

  return (
    <Box backdropFilter="blur(50px)" bgColor="rgba(255,255,255,0.1)" flex="1" borderRadius="10px" p={6} rounded="md">
      <Stack direction="row" spacing={2} mb={6} justifyContent="space-between">
        {Object.keys(filterCategories).map((filter) => (
          <Button
            key={filter}
            onClick={() => handleFilterClick(filter)}
            isActive={activeFilter === filter}
            _active={{
              transform: 'scale(0.95)',
            }}
            _hover={{
              bgColor: 'rgba(255,128,0, 0.25)',
              color: 'complimentary.300',
            }}
            color="white"
            minW={'180px'}
            colorScheme={activeFilter === filter ? 'orange' : 'gray'}
            variant={activeFilter === filter ? 'solid' : 'outline'}
          >
            {filter}
          </Button>
        ))}
      </Stack>
      <Box maxH={'480px'} overflow={'auto'}>
        <Table color={'white'} variant="simple">
          <Thead position="sticky">
            <Tr>
              <Th color={'complimentary.900'}>
                Asset Pair <ChevronDownIcon />
              </Th>
              <Th color={'complimentary.900'} isNumeric>
                APY <ChevronDownIcon />
              </Th>
              <Th color={'complimentary.900'} isNumeric>
                TVL <ChevronDownIcon />
              </Th>
              <Th color={'complimentary.900'}>Provider</Th>
              <Th color={'complimentary.900'}>Action</Th>
            </Tr>
          </Thead>
          <Tbody>
            {filteredData.map((asset, index) => (
              <Tr _even={{ bg: 'rgba(255, 128, 0, 0.1)' }} borderBottomColor={'transparent'}>
                <Td borderBottomColor="transparent">
                  <Flex align="center">
                    <Box w="2rem" h="2rem" bg="gray.200" rounded="full" mr={2}></Box>
                    <Text>{asset.assetPair}</Text>
                  </Flex>
                </Td>
                <Td borderBottom="0" borderBottomColor="transparent" isNumeric>
                  {formatApy(asset.apy)}
                </Td>
                <Td borderBottomColor="transparent" isNumeric>
                  {asset.tvl}
                </Td>
                <Td borderBottomColor="transparent">{asset.provider}</Td>
                <Td borderBottomColor="transparent">
                  <Button backgroundColor="rgba(255, 128, 0, 0.8)" rightIcon={<ExternalLinkIcon />} variant="ghost">
                    {actionTitles[asset.action.toLowerCase().replace(/\s+/g, '-') as keyof typeof actionTitles]}
                  </Button>
                </Td>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </Box>
    </Box>
  );
};

export default DefiTable;
