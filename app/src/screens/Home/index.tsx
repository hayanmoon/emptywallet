import {NativeStackNavigationProp} from '@react-navigation/native-stack';
import React from 'react';
import {Button, Text, View} from 'react-native';
import Layout from '../../Layout';

// type Props = NativeStackScreenProps<RootStackParamList, 'Home'>;
type HomeScreenNavigationProp = NativeStackNavigationProp<
  RootStackParamList,
  'Home'
>;

type HomeProps = {
  navigation: HomeScreenNavigationProp; //Props['navigation'];
};

function Home({navigation}: HomeProps) {
  const Logout = () => {
    navigation.reset({
      index: 0,
      routes: [{name: 'Login'}],
    });
  };
  return (
    <Layout>
      <View>
        <Text>Home</Text>
        <Button title="Logout" onPress={Logout}>
          Logout
        </Button>
      </View>
    </Layout>
  );
}

export default Home;
