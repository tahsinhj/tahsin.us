import React, { useState, useRef, useEffect } from 'react';
import { useSpring, animated } from 'react-spring';

export default function ScrollingText(props: { text: string }) {
  const parent = useRef(null);
  const text = useRef(null);
  const [shouldScroll, scroll] = useState(false);

  const scrolling = useSpring({
    from: { transform: 'translate(0%,0)' },
    to: { transform: 'translate(-60%,0)' },
    config: { duration: 5000 },

    reset: true,
    loop: true,
    delay: 1000,
    //reverse: key % 2 == 0,
  });

  useEffect(() => {
    const parentWidth = (parent.current as any).offsetWidth;
    const textWidth = (text.current as any).scrollWidth;
    if (textWidth > parentWidth) {
      scroll(true);
    }
  }, []);

  return (
    <div ref={parent}>
      {shouldScroll ? (
        <animated.p style={scrolling} ref={text}>
          {props.text}
        </animated.p>
      ) : (
        <p ref={text}>{props.text}</p>
      )}
    </div>
  );
}
