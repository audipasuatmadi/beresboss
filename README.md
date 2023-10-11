# Beres Bos!!!
This project was created out of personal bad experience with the repetitiveness of having to "run the same command in different locations sequentially," such as executing 'mvn clean install' in multiple locations. Why not simply define the command once, and then specify the locations where the projects are located? right?...

# Usage
To use this software, it's pretty easy. Add it to your user / system's `Path` variable, and then run the command as follows:
```bash
beresboss "cmd" ./dir1 ./dir2 ./parentdir1/dir3 /from/root/dir
```

A "more" real example:
```bash
beresboss "mvn clean install -T2C --settings \"/path/to/my-settings.xml\"" ./project1 ./module1/project2 ./module1/project3
```

# Demo Example
![Alt text](./docs/sample.gif)


# Contributing
Feel free to create a PR!

![Alt text](./docs/image-1.png)

# Contributors
* Putu Audi Pasuatmadi (audipasuatmadi@gmail.com)
