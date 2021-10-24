import React from 'react';
import {SafeAreaView, StyleSheet} from 'react-native';

type Props = {
  children: JSX.Element;
};

function Layout({children}: Props) {
  return <SafeAreaView style={styles.layout}>{children}</SafeAreaView>;
}

const styles = StyleSheet.create({
  layout: {
    margin: 32,
    paddingHorizontal: 24,
  },
});

export default Layout;
