import React from 'react';
import {Button, TextInput, View, Text, StyleSheet, Alert} from 'react-native';
import {useForm, Controller, SubmitHandler} from 'react-hook-form';
import {Auth} from 'aws-amplify';

import Layout from '../../Layout';
import {NativeStackScreenProps} from '@react-navigation/native-stack';

type FormData = {
  email: string;
  code: string;
};

type Props = NativeStackScreenProps<RootStackParamList, 'Confirm'>;

const Confirm = ({navigation}: Props) => {
  const {
    control,
    handleSubmit,
    formState: {errors},
  } = useForm<FormData>();

  const onSubmit: SubmitHandler<FormData> = async data => {
    await confirmSignUp(data);
    console.log(data, 'data');
  };

  async function confirmSignUp({email, code}: FormData) {
    try {
      await Auth.confirmSignUp(email, code);
      navigation.reset({
        index: 0,
        routes: [{name: 'Login'}],
      });
    } catch (error) {
      console.log('error confirming sign up', error);
    }
  }

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
              placeholder="email"
            />
          )}
          name="email"
          defaultValue=""
        />
        {errors.code && <Text>This is required.</Text>}
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
              placeholder="code"
            />
          )}
          name="code"
          defaultValue=""
        />
        {errors.code && <Text>This is required.</Text>}
        <Button title="Submit" onPress={handleSubmit(onSubmit)} />
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

export default Confirm;
