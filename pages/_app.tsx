import '../styles/globals.css';
import { Chain } from '@chain-registry/types';
import { Box, ChakraProvider } from '@chakra-ui/react';
import {
  quicksilverProtoRegistry,
  quicksilverAminoConverters,
  getSigningQuicksilverClientOptions,
  getSigningCosmosClientOptions,
} from '@chalabi/quicksilverjs';
import { Registry } from '@cosmjs/proto-signing';
import { SigningStargateClientOptions, AminoTypes } from '@cosmjs/stargate';
import { SignerOptions, WalletViewProps } from '@cosmos-kit/core';
import { wallets as cosmostationWallets } from '@cosmos-kit/cosmostation';
import { wallets as keplrWallets } from '@cosmos-kit/keplr';
import { wallets as leapWallets } from '@cosmos-kit/leap';
import { ChainProvider } from '@cosmos-kit/react';
import { QueryClientProvider, QueryClient } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { chains, assets } from 'chain-registry';
import { cosmosAminoConverters, cosmosProtoRegistry } from 'interchain-query';
import type { AppProps } from 'next/app';

import { Header, SideHeader } from '@/components';
import { defaultTheme } from '@/config';

import '@interchain-ui/react/styles';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      refetchOnWindowFocus: false,
    },
  },
});

function CreateCosmosApp({ Component, pageProps }: AppProps) {
  const signerOptions: SignerOptions = {
    signingStargate: (chain: Chain): SigningStargateClientOptions | undefined => {
      const registry = new Registry(cosmosProtoRegistry);
      const aminoTypes = new AminoTypes(cosmosAminoConverters);
      return {
        aminoTypes: aminoTypes,
        registry: registry,
      };
    },
  };

  return (
    <ChakraProvider theme={defaultTheme}>
      <ChainProvider
        chains={chains}
        assetLists={assets}
        wallets={[...keplrWallets, ...cosmostationWallets, ...leapWallets]}
        walletConnectOptions={{
          signClient: {
            projectId: 'a8510432ebb71e6948cfd6cde54b70f7',
            relayUrl: 'wss://relay.walletconnect.org',
            metadata: {
              name: 'CosmosKit Template',
              description: 'CosmosKit dapp template',
              url: 'https://docs.cosmoskit.com/',
              icons: [],
            },
          },
        }}
        signerOptions={signerOptions}
      >
        <QueryClientProvider client={queryClient}>
          <ReactQueryDevtools initialIsOpen={true} />
          <Header chainName="quicksilver" />
          <SideHeader />
          <Component {...pageProps} />
        </QueryClientProvider>
      </ChainProvider>
    </ChakraProvider>
  );
}

export default CreateCosmosApp;
