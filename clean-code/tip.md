### Make function calls readable

**BAD:**
```java
BoundingBox box = new BoundingBox(10, 50, 20, 5);
```

**GOOD:**
```java
BoundingBox box = new BoundingBox(
    /* top= */ 10,
    /* right= */ 50,
    /* bottom= */ 20,
    /* left= */ 5);
```

### Anonymous functions can be great for small things

**BAD:**
```java
List<Feedback> getUsefulFeedback(List<Feedback> allFeedback) {
  return allFeedback
      .filter(feedback -> !feedback.getComment().isEmpty());     ❶
}
```

**GOOD:**
```java
List<Feedback> getUsefulFeedback(List<Feedback> allFeedback) {
  return allFeedback.filter(hasNonEmptyComment);                ❶
}
 
private Boolean hasNonEmptyComment(Feedback feedback) {         ❷
  return !feedback.getComment().isEmpty();
}
```

### New features can improve code
**BAD:**
```java
List<String> getNonEmptyStrings(List<String> strings) {
  List<String> nonEmptyStrings = new ArrayList<>();
  for (String str : strings) {
    if (!str.isEmpty()) {
      nonEmptyStrings.add(str);
    }
  }
  return nonEmptyStrings;
}
```

**GOOD:**
```java
List<String> getNonEmptyStrings(List<String> strings) {
  return strings
      .stream()
      .filter(str -> !str.isEmpty())
      .collect(toList());
}
```