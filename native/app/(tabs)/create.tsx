import { Text, View, Pressable } from 'react-native';
import { styled } from 'nativewind';
import { LinearGradient } from 'expo-linear-gradient';
import { MaterialIcons } from '@expo/vector-icons';

const CustomText = styled(Text);
const CustomView = styled(View);
const GradientBackground = styled(LinearGradient);
const EditButton = styled(Pressable);

export default function CreateScreen() {
    return (
        <CustomView className="flex-1 items-center justify-center bg-background">
            <GradientBackground
                className="flex-1 w-full items-center justify-center"
                colors={['#99ffc8', '#c4effd']}
                start={{ x: 0, y: 0 }}
                end={{ x: 1, y: 1 }}
            >
                <CustomView
                    className="flex-1 bg-white w-5/6 my-8 rounded-2xl"
                    style={{ elevation: 7 }}
                >
                </CustomView>
                <Pressable
                    style={({ pressed }) => ({
                        backgroundColor: pressed ? '#8AE1B2' : '#99ffc8',
                        elevation: 7,
                        borderRadius: 50,
                        padding: 12,
                        position: 'absolute',
                        bottom: 20,
                        right: 20,
                    })}
                >
                    <MaterialIcons name="edit" size={36} color="black" />
                </Pressable>
            </GradientBackground>
        </CustomView>
    );
}
