import { View, Text } from 'react-native';
import { styled } from 'nativewind';
import { MaterialIcons } from '@expo/vector-icons';

const CustomText = styled(Text);
const CustomView = styled(View);

export default function CreateHeader() {
    return (
        <CustomView className="flex-1 flex-row justify-between h-24 w-full p-4 items-center">
            <CustomText className="font-hanken">Cancel</CustomText>
            <CustomView className="flex-row items-center space-x-2">
                <CustomText className="font-inter font-bold text-lg">Card Name</CustomText>
                <MaterialIcons style={{ lineHeight: 32 }} name="edit" size={16} color="black" />
            </CustomView>
            <CustomText className="font-hanken">Save</CustomText>
        </CustomView>
    )
}
