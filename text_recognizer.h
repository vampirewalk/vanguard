//
//  text_recognizer.h
//  STR
//
//  Created by vampirewalk on 2015/4/14.
//  Copyright (c) 2015年 mocacube. All rights reserved.
//

#ifndef __STR__text_recognizer__
#define __STR__text_recognizer__

#include <iostream>
#include <string>
#include <vector>

#include <opencv2/text.hpp>
#include <opencv2/core/utility.hpp>
#include <opencv2/highgui.hpp>
#include <opencv2/imgproc.hpp>

class TextRecognizer {
    
public:
    TextRecognizer(){};
    std::vector<std::string> recognize(std::vector<char> data);
};

#endif /* defined(__STR__text_recognizer__) */
