import React from 'react';
import {Button, TextInput, View, Text, StyleSheet} from 'react-native';
import {useForm, Controller, SubmitHandler} from 'react-hook-form';
import {NativeStackScreenProps} from '@react-navigation/native-stack';

import Layout from '../../Layout';
// import {RootStackParamList} from '../../../App';

type FormData = {
  email: string;
  password: string;
};

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const Login = ({navigation}: Props) => {
  const {
    control,
    handleSubmit,
    formState: {errors},
  } = useForm<FormData>();
  const onSubmit: SubmitHandler<FormData> = data => {
    console.log(data, 'data');
  };
  return (
    <Layout>
      <View>
        <Controller
          control={control}
          rules={{
            required: true,
          }}
          render={({field: {onChange, value}}) => (
            <TextInput
              style={styles.input}
              onChangeText={onChange}
              value={value}
              placeholder="username"
            />
          )}
          name="email"
          defaultValue=""
        />
        {errors.email && <Text>This is required.</Text>}
        <Controller
          control={control}
          rules={{
            required: true,
          }}
          render={({field: {onChange, value}}) => (
            <TextInput
              secureTextEntry={true}
              style={styles.input}
              onChangeText={onChange}
              value={value}
              placeholder="password"
            />
          )}
          name="password"
          defaultValue=""
        />
        {errors.password && <Text>This is required.</Text>}
        <Button title="Submit" onPress={handleSubmit(onSubmit)} />
        <Button
          title="Register"
          onPress={() => navigation.navigate('Register')}
        />
      </View>
    </Layout>
  );
};

const styles = StyleSheet.create({
  input: {
    height: 40,
    borderWidth: 1,
    marginBottom: 10,
    padding: 10,
  },
});

export default Login;
