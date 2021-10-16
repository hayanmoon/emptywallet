import React from 'react';
import {SafeAreaView, ScrollView, StyleSheet} from 'react-native';

type Props = {
  children: JSX.Element;
};

const Layout = ({children}: Props) => {
  return (
    <SafeAreaView style={styles.layout}>
      <ScrollView>{children}</ScrollView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  layout: {
    margin: 32,
    paddingHorizontal: 24,
  },
});

export default Layout;
